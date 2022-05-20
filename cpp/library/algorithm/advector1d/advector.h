#pragma once

#include <utility>

#include "base/field1d.h"

namespace advection {

class Advector {
  public:
	Advector() = default;
	virtual ~Advector() = default;

	virtual std::pair<Field1D*, Field1D*> Process(Field1D* field, Field1D* buff) = 0;
};

};  // advection
