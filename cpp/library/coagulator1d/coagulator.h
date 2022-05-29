#pragma once

#include "base/field1d.h"

namespace coagulation {

class Coagulator1D {
  public:
	Coagulator1D() = default;
	virtual ~Coagulator1D() = default;

	virtual std::pair<Field1D*, Field1D*> Process(Field1D* field, Field1D* buff) = 0;
};

};  // namespace coagulation
