#include "coagulator.h"

#include <memory>

namespace coagulation {

std::pair<Field1D*, Field1D*> SequentialCoagulator1D::Process(Field1D* field, Field1D* buff) {
	Field1D& f = *field;
	Field1D& b = *buff;

	for (size_t i = 0; i < f.Size(); i++) {
		base_coagulator_->Process(&f[i], &b[i], f.Volumes());
	}

	return std::make_pair(field, buff);
}

};  // namespace coagulation
