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
	void AssignSum(double coef1, const Cell& cell1, double coef2, const Cell& cell2, double coef3, const Cell& cell3);

  private:
	std::vector<double> data_;
};

