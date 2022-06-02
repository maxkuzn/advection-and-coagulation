#pragma once

#include "field1d.h"

#include <fstream>

class FieldSaver {
  public:
    FieldSaver() = default;

    virtual ~FieldSaver() = default;

    virtual void Save(const Field1D& field) = 0;
};
