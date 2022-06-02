#include "coagulator.h"

#include <memory>
#include <vector>
#include <thread>
#include <iostream>

namespace coagulation {

std::pair<Field1D*, Field1D*> ParallelPoolCoagulator1D::Process(Field1D* field, Field1D* buff) {
    std::vector<std::thread> threads;
    threads.reserve(field->Size());

    for (size_t i = 0; i < field->Size(); i += batch_size_) {
        size_t start = i;
        size_t end = std::min(i + batch_size_, field->Size());

        std::unique_lock lock(mtx_);

        queue_.push(std::make_tuple(field, buff, start, end));

        cv_.notify_one();
    }

    for (auto&& t: threads) {
        t.join();
    }

    return std::make_pair(field, buff);
}

void ParallelPoolCoagulator1D::DoWork() {
    for (;;) {
        std::unique_lock lock(mtx_);
        cv_.wait(lock, [&] {
            return stop_ || !queue_.empty();
        });

        if (stop_) {
            return;
        }

        auto [field, buff, start, end] = queue_.front();
        queue_.pop();

        lock.unlock();

        Field1D& f = *field;
        Field1D& b = *buff;


        for (size_t i = start; i < end; i++) {
            base_coagulator_->Process(&f[i], &b[i], volumes_);
        }
    }
}

};  // namespace coagulation
