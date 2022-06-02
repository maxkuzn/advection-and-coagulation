#include "coagulator.h"

#include <memory>
#include <vector>
#include <thread>

namespace coagulation {

std::pair<Field1D*, Field1D*> NaiveParallelCoagulator1D::Process(Field1D* field, Field1D* buff) {
    std::vector<std::thread> threads;
    threads.reserve(field->Size());

    for (size_t i = 0; i < field->Size(); i += batch_size_) {
        size_t start = i;
        size_t end = std::min(i + batch_size_, field->Size());

        threads.emplace_back(&NaiveParallelCoagulator1D::DoWork, this, field, buff, field->Volumes(), start, end);
    }

    for (auto&& t: threads) {
        t.join();
    }

    return std::make_pair(field, buff);
}

void NaiveParallelCoagulator1D::DoWork(Field1D* field, Field1D* buff, const std::vector<double>& volumes, size_t begin,
                                       size_t end) {
    Field1D& f = *field;
    Field1D& b = *buff;

    for (size_t i = begin; i < end; i++) {
        base_coagulator_->Process(&f[i], &b[i], volumes);
    }
}

} // namespace coagulation
