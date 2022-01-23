#include "forward_difference.h"


void forw_diff(const Field1D& field_in, Field1D& field_out, const double sigma) {
	size_t size = field_in.Size();

	field_out[size - 1].AssignSum(1+sigma, field_in[size - 1], -sigma, field_out[0]);
	for (size_t x = 0; x < size - 1; x++) {
		field_out[x].AssignSum(1+sigma, field_in[x], -sigma, field_in[x + 1]);
	}
}

