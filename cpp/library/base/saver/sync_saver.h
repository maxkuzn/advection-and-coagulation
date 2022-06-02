#pragma once

#include "base/field1d.h"
#include "base/field1d_saver.h"

#include <fstream>

class SyncFieldSaver : public FieldSaver {
  public:
    SyncFieldSaver(const std::string_view& filename)
            : file_(filename), first_save_(true) {
    }

    ~SyncFieldSaver() = default;

    void Save(const Field1D& field) override {
        if (first_save_) {
            first_save_ = false;
            file_ << field.Size() << '\n';
        }
        for (size_t i = 0; i < field.Size(); i++) {
            auto& cell = field[i];
            for (size_t j = 0; j < cell.Size(); j++) {
                file_ << cell[j];
                if (j + 1 != cell.Size()) {
                    file_ << ' ';
                } else {
                    file_ << '\n';
                }
            }
        }
    }

  private:
    std::ofstream file_;
    bool first_save_;
};
