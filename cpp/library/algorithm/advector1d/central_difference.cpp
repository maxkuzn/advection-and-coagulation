#include "central_difference.h"

namespace advection {

std::pair<Field1D*, Field1D*> CentralDifference::Process(Field1D* field, Field1D* buff) {
	Field1D& field_in = *field;
	Field1D& field_out = *buff;

	auto s = field_in.Size();

	field_out[0].AssignSum(
			1, field_in[0],
			-sigma_/2, field_in[1],
			0, field_in[s-1]
	);

	field_out[s-1].AssignSum(
			1, field_in[s-1],
			0, field_in[0],
			sigma_/2, field_in[s-2]
	);

	for (size_t x = 1; x < s-1; x++) {
		field_out[x].AssignSum(
				1, field_in[x],
				-sigma_/2, field_in[x+1],
				sigma_/2, field_in[x-1]
		);
	}


	return std::make_pair(&field_out, &field_in);
}

};  // namespace advection
