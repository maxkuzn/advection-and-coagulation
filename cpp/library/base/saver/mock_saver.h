#pragma once

#include "base/field1d.h"
#include "base/field1d_saver.h"

#include <fstream>

class MockFieldSaver : public FieldSaver {
  public:
    MockFieldSaver() = default;

    ~MockFieldSaver() = default;

    void Save(const Field1D& /*field*/) override {
    }
};
