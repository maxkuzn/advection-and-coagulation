#pragma once

#include <vector>


class Cell {
  public:
	Cell(size_t particle_sizes_num)
		: data_(particle_sizes_num, 0.0)
	{
	}

	double& operator[](size_t idx) {
		return data_[idx];
	}

	double operator[](size_t idx) const {
		return data_[idx];
	}

	size_t Size() const {
		return data_.size();
	}

	void AssignSum(double coef1, const Cell& cell1, double coef2, const Cell& cell2);

  private:
	std::vector<double> data_;
};

class Field1D {
  public:
	Field1D(size_t field_size, size_t particle_sizes_num)
		: data_(field_size, Cell(particle_sizes_num))
	{
	}

	Cell& operator[](size_t idx) {
		return data_[idx];
	}

	const Cell& operator[](size_t idx) const {
		return data_[idx];
	}

	size_t Size() const {
		return data_.size();
	}

  private:
	std::vector<Cell> data_;
};
