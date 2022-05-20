#include "coagulator.h"

namespace coagulation {

double PredCorrCoagulator::Process(Cell* cell, Cell* buff, const std::vector<double>& volumes) {
	(void) cell;
	(void) buff;
	(void) volumes;
	(void) kernel_;
	(void) time_step_;

	return 0;
}

};  // namespace coagulation
