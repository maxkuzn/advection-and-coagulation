#pragma once

#include "algorithm/coagulation/coagulator.h"
#include "algorithm/coagulation/kernel.h"

#include <vector>
#include <memory>


namespace coagulation {

class PredCorrCoagulator : public Coagulator {
  public:
	PredCorrCoagulator(std::shared_ptr<Kernel> kernel, double time_step)
		: kernel_(kernel)
		, time_step_(time_step)
	{}

	~PredCorrCoagulator() = default;

	void Process(Cell* cell, Cell* buff, const std::vector<double>& volumes) override;

  private:
	double ProcessHalf(Cell& cell, const std::vector<double>& volumes, size_t idx);
	double ProcessFull(Cell& cell, const std::vector<double>& volumes, size_t idx);
	double ComputeL1(Cell& cell, const std::vector<double>& volumes, size_t idx);
	double ComputeL2(Cell& cell, const std::vector<double>& volumes, size_t idx);

	std::shared_ptr<Kernel> kernel_;
	const double time_step_;
};

};  // namespace coagulation
