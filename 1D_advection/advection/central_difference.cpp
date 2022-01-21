#include "forward_difference.h"


std::vector<std::vector<double>> cent_diff(
		const size_t T,
		const double sigma,
		const std::vector<double>& u
) {
	size_t size = u.size();
	std::vector<std::vector<double>> history(
			T + 1,
			std::vector<double>(size)
	);

	history[0] = u;

	for (size_t t = 0; t < T; t++) {
		std::vector<double>& u_prev = history[t];
		std::vector<double>& u_next = history[t + 1];

		u_next.back() = u_prev.back() - sigma / 2 * u_prev[0] + sigma / 2 * u_prev[size - 2];
		u_next[0] = u_prev[0] - sigma / 2 * u_prev[1] + sigma / 2 * u_prev[size - 1];
		for (size_t x = 1; x < size - 1; x++) {
			u_next[x] = u_prev[x] - sigma / 2 * u_prev[x + 1] + sigma / 2 * u_prev[x - 1];
		}
	}

	return history;
}

