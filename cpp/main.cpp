#include "config/config.h"

#include "base/init_field1d.h"
#include "base/field1d_saver.h"

#include "algorithm/advector1d/advector.h"
#include "algorithm/advector1d/central_difference.h"
#include "algorithm/coagulation/predcorr/coagulator.h"
#include "algorithm/coagulation/fast/coagulator.h"
#include "algorithm/coagulation/kernel/identity.h"
#include "algorithm/coagulation/kernel/addition.h"

#include "coagulator1d/sequential/coagulator.h"
#include "coagulator1d/naiveparallel/coagulator.h"
#include "coagulator1d/parallelpool/coagulator.h"

#include "util/progress.h"

#include <cstdlib>
#include <iostream>
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

std::shared_ptr<coagulation::Coagulator1D> chooseCoagulator(const Config& cfg, const std::vector<double>& volumes) {
    std::shared_ptr<coagulation::Kernel> kernel;
    if (cfg.coagulation_kernel_name == "Identity") {
        kernel = std::shared_ptr<coagulation::Kernel>(
                new coagulation::IdentityKernel()
        );
    } else if (cfg.base_coagulator_name == "Addition") {
        kernel = std::shared_ptr<coagulation::Kernel>(
                new coagulation::AdditionKernel()
        );
    } else {
        throw std::runtime_error("Unknown coagulation kernel");
    }

    std::shared_ptr<coagulation::Coagulator> base_coagulator;
    if (cfg.base_coagulator_name == "PredictorCorrector") {
        base_coagulator = std::shared_ptr<coagulation::Coagulator>(
                new coagulation::PredCorrCoagulator(kernel, cfg.time_step)
        );
    } else if (cfg.base_coagulator_name == "Fast") {
        base_coagulator = std::shared_ptr<coagulation::Coagulator>(
                new coagulation::FastCoagulator(kernel, cfg.time_step, volumes)
        );
    } else {
        throw std::runtime_error("Unknown base coagulator");
    }

    std::shared_ptr<coagulation::Coagulator1D> coagulator;
    if (cfg.coagulator_name == "Sequential") {
        coagulator = std::shared_ptr<coagulation::Coagulator1D>(
                new coagulation::SequentialCoagulator1D(base_coagulator)
        );
    } else if (cfg.coagulator_name == "NaiveParallel") {
        coagulator = std::shared_ptr<coagulation::Coagulator1D>(
                new coagulation::NaiveParallelCoagulator1D(base_coagulator)
        );
    } else if (cfg.coagulator_name == "ParallelPool") {
        coagulator = std::shared_ptr<coagulation::Coagulator1D>(
                new coagulation::ParallelPoolCoagulator1D(base_coagulator, volumes)
        );
    } else {
        throw std::runtime_error("Unknown coagulator");
    }

    return coagulator;
}

int main() {
    // 1
    Config cfg{
            .field_size = 1.0,
            .field_cells_size = 200,
            .particles_sizes_num = 200,
            .min_particle_size = 0.1,
            .max_particle_size = 1.0,

            .total_time = 20.0,
            .time_steps = 500,

            .advection_coef = 0.1,

            .advector_name = "CentralDifference",
            .coagulation_kernel_name = "Identity", // "Identity", "Addition"
            .base_coagulator_name = "Fast", // "PredictorCorrector", "Fast"
            .coagulator_name = "Sequential",  // "Sequential", "NaiveParallel", "ParallelPool"
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

    auto coagulator = chooseCoagulator(cfg, field->Volumes());

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
