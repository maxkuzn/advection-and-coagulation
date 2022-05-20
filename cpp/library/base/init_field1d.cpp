#include "init_field1d.h"

#include <cmath>

namespace {

double gaussian_pdf(double mu, double sigma, double x) {
	constexpr double inv_sqrt_2pi = 0.3989422804014327;
	double y = (x - mu) / sigma;

	return inv_sqrt_2pi / sigma * std::exp(-0.5 * y * y);
}

double size_factor(double vMin, double v) {
	constexpr double sigma_2 = 0.1;
	double sigma = std::sqrt(sigma_2);

	double minF = gaussian_pdf(vMin, sigma, vMin);
	double f = gaussian_pdf(vMin, sigma, v);

	return f / minF;
}

double coord_factor(size_t x, size_t limit) {
	if (x >= limit) {
		return 0.0;
	}
	double y = 2 * M_PI * x / limit;
	return (std::cos(y - M_PI) + 1) / 2;
}

};

Field1D init_field_1d(size_t field_size, size_t particle_sizes_num, double vMin, double vMax) {
	Field1D field(field_size, particle_sizes_num, vMin, vMax);

	// Fill only first 25% of the field with factor func
	size_t limit = field_size / 4;
	for (size_t x = 0; x < limit; x++) {
		double factor = ::coord_factor(x, limit);

		for (size_t i = 0; i < particle_sizes_num; i++) {
			double v = field.Volume(i);
			field[x][i] = factor * ::size_factor(vMin, v);
		}
	}

	return field;
}
