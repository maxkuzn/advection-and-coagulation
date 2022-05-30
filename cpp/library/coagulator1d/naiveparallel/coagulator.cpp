#include "coagulator.h"

#include <memory>
#include <vector>
#include <thread>

namespace coagulation {

std::pair<Field1D*, Field1D*> NaiveParallelCoagulator1D::Process(Field1D* field, Field1D* buff) {
    Field1D& f = *field;
    Field1D& b = *buff;

    std::vector<std::thread> threads;
    threads.reserve(f.Size());

    for (size_t i = 0; i < f.Size(); i++) {
        /*
        threads.emplace_back([&] () {
            base_coagulator_->Process(&f[i], &b[i], field->Volumes());
        });
         */

        threads.emplace_back(&NaiveParallelCoagulator1D::DoWork, this, &f[i], &b[i], field->Volumes());
    }

    for (auto&& t: threads) {
        t.join();
    }

    return std::make_pair(field, buff);
}

void NaiveParallelCoagulator1D::DoWork(Cell* cell, Cell* buff, const std::vector<double>& volumes) {
    base_coagulator_->Process(cell, buff, volumes);
}

} // namespace coagulation
