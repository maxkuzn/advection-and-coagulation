#pragma once

#include "base/field1d.h"
#include "coagulation_1D/kernel/base.h"

class Coagulator {
  public:
	Coagulator(const BaseKernel& kernel, double size_step, double time_step)
		: kernel_(kernel)
		, size_step_(size_step)
		, time_step_(time_step)
	{}

	void Process(Field1D& field, Field1D& field_buf);

  private:
	void ProcessCell(Cell& cell, Cell& cell_buf);

	double ProcessSizeHalf(const Cell& cell, size_t size_idx);
	double ProcessSizeFull(const Cell& cell, size_t size_idx);

	double ComputeL1(const Cell& cell, size_t size_idx);
	double ComputeL2(const Cell& cell, size_t size_idx);

  private:
	const BaseKernel& kernel_;
	double size_step_;
	double time_step_;
};

