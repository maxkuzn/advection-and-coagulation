#include "forward_difference.h"


std::vector<std::vector<double>> forw_diff(
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

		u_next.back() = (1 + sigma) * u_prev.back() - sigma * u_prev[0];
		for (size_t x = 0; x < size - 1; x++) {
			u_next[x] = (1 + sigma) * u_prev[x] - sigma * u_prev[x + 1];
		}
	}

	return history;
}

