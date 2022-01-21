#include "util.h"

#include <fstream>

void save_history(const std::vector<std::vector<double>>& history) {
	std::ofstream file("history.txt");

	size_t T = history.size();
	size_t N = history[0].size();

	file << T << ' ' << N << '\n';
	for (size_t t = 0; t < T; t++) {
		for (size_t x = 0; x < N; x++) {
			file << history[t][x];
			if (x + 1 == N) {
				file << '\n';
			} else {
				file << ' ';
			}
		}
	}
}

