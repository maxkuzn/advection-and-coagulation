#include "advection_1D/forward_difference.h"
#include "advection_1D/backward_difference.h"
#include "advection_1D/central_difference.h"
#include "advection_1D/util.h"
#include "base/init_field_1d.h"
#include "base/field_saver_1d.h"

#include <math.h>
#include <string_view>

constexpr std::string_view kHistoryFilename = "data/history.txt";

int main() {
	size_t field_size = 500;
	size_t particle_sizes_num = 100;

	// size_t time_steps = 500;
	// double advection_coef = 0.01;

	auto field = init_field_1d(field_size, particle_sizes_num);
	FieldSaver saver(kHistoryFilename);
	saver.Save(field);
	saver.Save(field);
	saver.Save(field);
	saver.Save(field);

	// auto history = cent_diff(T, sigma, u);
	// save_history(kHistoryFilename, history);
}
