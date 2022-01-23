#include "advection_1D/forward_difference.h"
#include "advection_1D/backward_difference.h"
#include "advection_1D/central_difference.h"
#include "base/init_field_1d.h"
#include "base/field_saver_1d.h"

#include <math.h>
#include <string_view>

constexpr std::string_view kHistoryFilename = "data/history.txt";

int main() {
	size_t field_size = 100;
	size_t particle_sizes_num = 10;

	size_t time_steps = 500;
	double advection_coef = 0.01;

	FieldSaver saver(kHistoryFilename);

	auto field1 = init_field_1d(field_size, particle_sizes_num);
	Field1D field2(field_size, particle_sizes_num);

	Field1D* field = &field1;
	Field1D* field_buf = &field2;

	for (size_t t = 0; t < time_steps; t++) {
		saver.Save(*field);
		forw_diff(*field, *field_buf, advection_coef);
		std::swap(field, field_buf);
	}
	saver.Save(*field);
}
