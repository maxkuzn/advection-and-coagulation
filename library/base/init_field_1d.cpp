#include "init_field_1d.h"

#include <cmath>

double gaussian_pdf(double x) {
	constexpr double sigma_2 = 0.3;
	constexpr double mu = 0.0;

	return 1.0 / std::sqrt(2.0 * sigma_2 * M_PI) * std::exp(
		-1.0 / 2.0 * (x - mu) * (x - mu) / sigma_2
	);
}

double coord_factor(size_t x, size_t limit) {
	if (x >= limit) {
		return 0.0;
	}
	double y = 2 * M_PI * x / limit;
	return (std::cos(y - M_PI) + 1) / 2;
}

Field1D init_field_1d(size_t field_size, size_t particle_sizes_num) {
	constexpr double Vmin = 0.1;
	constexpr double Vmax = 1.0;

	Field1D field(field_size, particle_sizes_num);

	// Fill only first 10% of the field with factor func
	size_t limit = field_size / 10;
	for (size_t x = 0; x < field_size; x++) {
		double factor = coord_factor(x, limit);

		for (size_t size_idx = 0; size_idx < particle_sizes_num; size_idx++) {
			double v = Vmin + (Vmax - Vmin) / (particle_sizes_num - 1) * size_idx;
			field[x][size_idx] = factor * gaussian_pdf(v);
		}
	}

	return field;
}

