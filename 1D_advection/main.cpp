#include "advection/forward_difference.h"
#include "advection/backward_difference.h"
#include "advection/central_difference.h"
#include "advection/util.h"

#include <math.h>

int main() {
	size_t N = 200;
	size_t T = 500;
	double sigma = 0.25;

	double h = 1.0 / (N - 1);
	std::vector<double> x(N);
	std::vector<double> u(N);

	for (size_t i = 0; i < N; i++) {
		x[i] = i * h;
		u[i] = sin(2*M_PI*x[i]);
	}

	auto history = cent_diff(T, sigma, u);
	save_history(history);
}
