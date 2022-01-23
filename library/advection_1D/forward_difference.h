#pragma once

#include <vector>

std::vector<std::vector<double>> forw_diff(
		const size_t T,
		const double sigma,
		const std::vector<double>& u);

