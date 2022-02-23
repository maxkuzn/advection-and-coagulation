#include "base/init_field_1d.h"
#include "base/field_saver_1d.h"

#include "advection_1D/forward_difference.h"
#include "advection_1D/backward_difference.h"
#include "advection_1D/central_difference.h"

#include "coagulation_1D/coagulation.h"
#include "coagulation_1D/kernel/identity.h"

#include "util/progress.h"

#include <math.h>
#include <string_view>
#include <iostream>

constexpr std::string_view kHistoryFilename = "data/history.txt";

int main() {
	size_t field_size = 100;
	size_t particle_sizes_num = 100;

	size_t time_steps = 500;
	double advection_coef = 0.1;

	double time_step = 10.0 / time_steps;
	double size_step = (1.0 - 0.1) / (particle_sizes_num - 1); // TODO: get that from Field

	IdentityKernel kernel;
	Coagulator coagulator(kernel, size_step, time_step);

	FieldSaver saver(kHistoryFilename);

	auto field1 = init_field_1d(field_size, particle_sizes_num);
	Field1D field2(field_size, particle_sizes_num);

	Field1D* field = &field1;
	Field1D* field_buf = &field2;

	std::cout << progress(0, time_steps) << '\r';
	std::cout.flush();

	saver.Save(*field);
	for (size_t t = 0; t < time_steps; t++) {
		cent_diff(*field, *field_buf, advection_coef);
		std::swap(field, field_buf);

		coagulator.Process(*field, *field_buf);

		saver.Save(*field);

		std::cout << progress(t + 1, time_steps) << '\r';
		std::cout.flush();
	}
}
