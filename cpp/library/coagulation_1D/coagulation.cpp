#include "coagulation.h"


void Coagulator::Process(Field1D& field, Field1D& field_buf) {
	for (size_t i = 0; i != field.Size(); i++) {
		ProcessCell(field[i], field_buf[i]);
	}
}

void Coagulator::ProcessCell(Cell& cell, Cell& cell_buf) {
	size_t N = cell.Size();
	for (size_t size_idx = 0; size_idx < N; size_idx++) {
		cell_buf[size_idx] = ProcessSizeHalf(cell, size_idx);
	}

	for (size_t size_idx = 0; size_idx < N; size_idx++) {
		cell[size_idx] = ProcessSizeFull(cell_buf, size_idx);
	}
}


double Coagulator::ProcessSizeHalf(const Cell& cell, size_t size_idx) {
	double L1 = ComputeL1(cell, size_idx);
	double L2 = ComputeL2(cell, size_idx);
	double v = cell[size_idx];

	return time_step_ / 2 * (L1 - v * L2) + v;
}

double Coagulator::ProcessSizeFull(const Cell& cell, size_t size_idx) {
	double L1 = ComputeL1(cell, size_idx);
	double L2 = ComputeL2(cell, size_idx);
	double v = cell[size_idx];

	return time_step_ * (L1 - v * L2) + v;
}


double Coagulator::ComputeL1(const Cell& cell_in, size_t size_idx) {
	if (size_idx == 0) {
		return 0.0;
	}

	double res = 0.0;
	for (size_t i = 0; i <= size_idx; i++) {
		size_t size1 = i;
		size_t size2 = size_idx - i;
		double add = kernel_.Compute(size1, size2) * cell_in[size1] * cell_in[size2];
		if (size1 == 0 || size2 == 0) {
			add /= 2;
		}
		res += add;
	}
	res *= size_step_ / 2.0;
	return res;
}

double Coagulator::ComputeL2(const Cell& cell, size_t size_idx) {
	double res = 0.0;
	res += size_step_ / 2 * kernel_.Compute(size_idx, 0.0) * cell[0];
	for (size_t i = 0; i < cell.Size(); i++) {
		double add = size_step_ * kernel_.Compute(size_idx, i) * cell[i];
		if (i == 0 || i + 1 == cell.Size()) {
			add /= 2;
		}
		res += add;
	}
	return res;
}

