#pragma once

#include "algorithm/coagulation/coagulator.h"
#include "base/field1d.h"
#include "coagulator1d/coagulator.h"

#include <memory>

namespace coagulation {

class SequentialCoagulator1D : public Coagulator1D {
  public:
    SequentialCoagulator1D(std::shared_ptr<Coagulator> base_coagulator)
            : base_coagulator_(base_coagulator) {}

    ~SequentialCoagulator1D() override = default;

    std::pair<Field1D*, Field1D*> Process(Field1D* field, Field1D* buff) override;

  private:
    std::shared_ptr<Coagulator> base_coagulator_;
};

};  // namespace coagulation
