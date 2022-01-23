#pragma once

#include <vector>
#include <string_view>

void save_history(const std::string_view& filename,
				  const std::vector<std::vector<double>>& history);

