#include "field_1d.h"

void Cell::AssignSum(double coef1, const Cell& cell1, double coef2, const Cell& cell2) {
	for (size_t i = 0; i < data_.size(); i++) {
		data_[i] = coef1 * cell1[i] + coef2 * cell2[i];
	}
}

