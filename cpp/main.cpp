#include "config/config.h"

#include "base/init_field1d.h"
#include "base/field1d_saver.h"

#include "algorithm/advector1d/advector.h"
#include "algorithm/advector1d/central_difference.h"
#include "algorithm/coagulation/predcorr/coagulator.h"
#include "algorithm/coagulation/kernel/identity.h"

#include "coagulator1d/sequential/coagulator.h"

#include "util/progress.h"

#include <cstdlib>
#include <iostream>
#include <math.h>
#include <memory>
#include <string_view>


constexpr std::string_view kHistoryFilename = "data/history.txt";


void run(
	const Config& cfg,
	Field1D* field, Field1D* buff,
	FieldSaver& saver,
	advection::Advector& advector,
	coagulation::Coagulator1D& coagulator
);

int main() {
	// 1
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
	auto kernel = std::shared_ptr<coagulation::Kernel>(
			new coagulation::IdentityKernel()
	);

	auto base_coagulator = std::shared_ptr<coagulation::Coagulator>(
			new coagulation::PredCorrCoagulator(kernel, cfg.time_step)
	);

	auto coagulator = std::shared_ptr<coagulation::Coagulator1D>(
			new coagulation::SequentialCoagulator1D(base_coagulator)
	);

	// 6
	// run
	run(
		cfg,
		field, field_buff,
		saver,
		*advector,
		*coagulator
	);

}

void run(
	const Config& cfg,
	Field1D* field, Field1D* buff,
	FieldSaver& saver,
	advection::Advector& advector,
	coagulation::Coagulator1D& coagulator
) {
	std::cout << progress(0, cfg.time_steps) << '\r';
	std::cout.flush();

	saver.Save(*field);
	for (size_t t = 0; t < cfg.time_steps; t++) {
		{
			auto [f, b] = advector.Process(field, buff);
			field = f;
			buff = b;
		}

		{
			auto [f, b] = coagulator.Process(field, buff);
			field = f;
			buff = b;
		}

		saver.Save(*field);

		std::cout << progress(t + 1, cfg.time_steps) << '\r';
		std::cout.flush();
	}
}
