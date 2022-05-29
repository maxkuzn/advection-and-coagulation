#pragma once

#include "base/cell.h"

#include <vector>


namespace coagulation {

class Coagulator {
  public:
	Coagulator() = default;
	virtual ~Coagulator() = default;

	virtual void Process(Cell* cell, Cell* buff, const std::vector<double>& volumes) = 0;
};

};  // namespace coagulation
