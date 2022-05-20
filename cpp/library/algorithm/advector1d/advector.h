#pragma once

class Advector {
  public:
	Advector() = default;
	virtual ~Advector() = default;

	virtual void Process(Field1D* field, field1D* buff) = 0;
};
