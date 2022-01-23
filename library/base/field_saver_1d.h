#pragma once

#include "field_1d.h"

#include <fstream>

class FieldSaver {
  public:
	FieldSaver(const std::string_view& filename)
		: file_(filename)
		, first_save_(true)
	{
	}

	void Save(const Field1D& field) {
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
