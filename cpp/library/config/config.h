#pragma once

#include <string>


struct Config {
    double field_size;
    size_t field_cells_size;
    size_t particles_sizes_num;
    double min_particle_size;
    double max_particle_size;

    double total_time;
    size_t time_steps;
    double time_step;

    double advection_coef;

    std::string saver_name;

    std::string advector_name;

    std::string coagulation_kernel_name;
    std::string base_coagulator_name;
    std::string coagulator_name;

    size_t batch_size;

    bool ValidateAndFill() {
        if (field_size <= 0) {
            return false;
        }

        if (field_cells_size <= 0) {
            return false;
        }

        if (particles_sizes_num <= 0) {
            return false;
        }

        if (min_particle_size <= 0) {
            return false;
        }

        if (max_particle_size <= 0) {
            return false;
        }

        if (total_time <= 0) {
            return false;
        }

        if (time_steps <= 0) {
            return false;
        }

        if (saver_name.empty()) {
            return false;
        }

        if (advector_name.empty()) {
            return false;
        }

        if (coagulation_kernel_name.empty()) {
            return false;
        }

        if (base_coagulator_name.empty()) {
            return false;
        }

        if (coagulator_name.empty()) {
            return false;
        }

        if (batch_size <= 0) {
            return false;
        }

        time_step = total_time / time_steps;
        return true;
    }
};
