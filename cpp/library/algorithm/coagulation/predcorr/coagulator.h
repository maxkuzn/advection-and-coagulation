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

	double Process(Cell* cell, Cell* buff, const std::vector<double>& volumes) override;

  private:
	std::shared_ptr<Kernel> kernel_;
	const double time_step_;
};

};  // namespace coagulation
