#pragma once

#include "advector.h"

namespace advection {

class CentralDifference : public Advector {
  public:
	CentralDifference(double sigma) : sigma_(sigma) {
	}

	~CentralDifference() = default;

	std::pair<Field1D*, Field1D*> Process(Field1D* field, Field1D* buff) override;

  private:
	const double sigma_;
};

};  // namespace advection
