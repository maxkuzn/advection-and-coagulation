#include "central_difference.h"

void cent_diff(const Field1D& field_in, Field1D& field_out, const double sigma) {
	size_t size = field_in.Size();

	field_out[0].AssignSum(
			       1, field_in[0],
			-sigma/2, field_in[1],
			 sigma/2, field_in[size-1]
	);
	field_out[size-1].AssignSum(
			       1, field_in[size-1],
			-sigma/2, field_in[0],
			 sigma/2, field_in[size-2]
	);
	for (size_t x = 1; x < size - 1; x++) {
		field_out[x].AssignSum(
				       1, field_in[x],
				-sigma/2, field_in[x+1],
			 	 sigma/2, field_in[x-1]
		);
	}
}

