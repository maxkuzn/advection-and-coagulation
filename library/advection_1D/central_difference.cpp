#include "forward_difference.h"


std::vector<std::vector<double>> cent_diff(
		const size_t T,
		const double sigma,
		const std::vector<double>& u
) {
	size_t size = u.size();
	std::vector<std::vector<double>> history(
			T + 1,
			std::vector<double>(size + 3)
	);

	for (size_t i = 0; i < size; i++) {
		history[0][i + 2] = u[i];
	}
	history[0][0] = u[size - 2];
	history[0][1] = u[size - 1];
	history[0].back() = u[0];

	for (size_t t = 0; t < T; t++) {
		std::vector<double>& u_prev = history[t];
		std::vector<double>& u_next = history[t + 1];

		for (size_t x = 2; x < u_next.size() - 1; x++) {
			// u_next[x] = u_prev[x] - sigma / 2 * u_prev[x + 1] + sigma / 2 * u_prev[x - 1];

			double sub = sigma / 2 * u_prev[x + 1];
			sub = std::min(sub, u_prev[x]);

			double add = sigma / 2 * u_prev[x - 1];
			add = std::min(add, u_prev[x - 2]);

			u_next[x] = u_prev[x] - sub + sigma / 2 * u_prev[x - 1];
		}

		u_next[0] = u_next[u_next.size() - 3];
		u_next[1] = u_next[u_next.size() - 2];
		u_next.back() = u_next[2];
	}

	// Erase first two and last elements (dublicates)
	for (size_t i = 0; i != history.size(); i++) {
		history[i].pop_back();
		history[i].erase(history[i].begin(), history[i].begin() + 1);
	}

	return history;
}

