#pragma once

#include "algorithm/coagulation/coagulator.h"
#include "base/field1d.h"
#include "coagulator1d/coagulator.h"

#include <memory>

namespace coagulation {

class NaiveParallelCoagulator1D : public Coagulator1D {
  public:
    explicit NaiveParallelCoagulator1D(std::shared_ptr<Coagulator> base_coagulator)
            : base_coagulator_(std::move(base_coagulator)) {}

    ~NaiveParallelCoagulator1D() override = default;

    std::pair<Field1D*, Field1D*> Process(Field1D* field, Field1D* buff) override;

  private:
    void DoWork(Cell* cell, Cell* buff, const std::vector<double>& volumes);

    std::shared_ptr<Coagulator> base_coagulator_;
};

}  // namespace coagulation
