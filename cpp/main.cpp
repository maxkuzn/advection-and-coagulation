#include "config/config.h"

#include "base/init_field1d.h"
#include "base/field1d_saver.h"

#include "algorithm/advector1d/advector.h"
#include "algorithm/advector1d/central_difference.h"

#include "util/progress.h"

#include <cstdlib>
#include <iostream>
#include <math.h>
#include <memory>
#include <string_view>



#include "coagulation_1D/coagulation.h"
#include "coagulation_1D/kernel/identity.h"


constexpr std::string_view kHistoryFilename = "data/history.txt";


void run(
	const Config& cfg,
	Field1D* field, Field1D* buff,
	FieldSaver& saver,
	advection::Advector& advector
	// Coagulator1D& coagulator
);

int main() {
	Config cfg{
		.field_size = 200,
		.field_cells_size = 200,
		.particles_sizes_num = 200,
		.min_particle_size = 0.1,
		.max_particle_size = 1.0,

		.total_time = 10.0,
		.time_steps = 1000,

		.advection_coef = 0.1,

		.advector_name = "CentralDifference",
		.coagulator_name = "Sequential",
		.coagulation_kernel_name = "Identity",
	};
	if (!cfg.ValidateAndFill()) {
		std::cerr << "Invalid config\n";
		std::exit(1);
	}

	size_t particle_sizes_num = 200;

	double size_step = (1.0 - 0.1) / (particle_sizes_num - 1); // TODO: get that from Field

	// 1
	// Config cfg;

	// 2
	FieldSaver saver(kHistoryFilename);

	// 3
	auto field1 = init_field_1d(
			cfg.field_cells_size,
			cfg.particles_sizes_num,
			cfg.min_particle_size, cfg.max_particle_size
	);
	Field1D field2(
			cfg.field_cells_size,
			cfg.particles_sizes_num,
			cfg.min_particle_size, cfg.max_particle_size
	);

	Field1D* field = &field1;
	Field1D* field_buff = &field2;

	// 4
	// advector
	auto advector = std::shared_ptr<advection::Advector>(
			new advection::CentralDifference(cfg.advection_coef)
	);

	// 5
	// coagulator
	IdentityKernel kernel;
	Coagulator coagulator(kernel, size_step, cfg.time_step);

	// 6
	// run
	run(
		cfg,
		field, field_buff,
		saver,
		*advector
		// coagulator
	);

}

void run(
	const Config& cfg,
	Field1D* field, Field1D* buff,
	FieldSaver& saver,
	advection::Advector& advector
	// Coagulator1D* coagulator
) {
	(void) cfg;
	(void) field;
	(void) buff;
	(void) saver;
	(void) advector;
	/*
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
	*/
}
