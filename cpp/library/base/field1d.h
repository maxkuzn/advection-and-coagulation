#pragma once

#include "cell.h"

#include <vector>


class Field1D {
  public:
	Field1D(size_t field_size, size_t particle_sizes_num, double vMin, double vMax)
		: cells_(field_size, Cell(particle_sizes_num))
	{
		volumes_.resize(particle_sizes_num);

		for (size_t i = 0; i < volumes_.size(); i++) {
			volumes_[i] = vMin + (vMax-vMin)*double(i)/double(volumes_.size()-1);
		}
	}

	Cell& operator[](size_t idx) {
		return cells_[idx];
	}

	const Cell& operator[](size_t idx) const {
		return cells_[idx];
	}

	size_t Size() const {
		return cells_.size();
	}

	const std::vector<double>& Volumes() const {
		return volumes_;
	}

	double Volume(size_t idx) const {
		return volumes_[idx];
	}

  private:
	std::vector<Cell> cells_;
	std::vector<double> volumes_;
};
