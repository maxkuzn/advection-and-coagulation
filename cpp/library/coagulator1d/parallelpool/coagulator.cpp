#include "coagulator.h"

#include <memory>
#include <vector>
#include <thread>
#include <iostream>

namespace coagulation {

std::pair<Field1D*, Field1D*> ParallelPoolCoagulator1D::Process(Field1D* field, Field1D* buff) {
    Field1D& f = *field;
    Field1D& b = *buff;

    std::vector<std::thread> threads;
    threads.reserve(f.Size());

    for (size_t i = 0; i < f.Size(); i++) {
        std::unique_lock lock(mtx_);

        queue_.push(std::make_pair(&f[i], &b[i]));

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

        auto [cell, buff] = queue_.front();
        queue_.pop();

        lock.unlock();

        base_coagulator_->Process(cell, buff, volumes_);
    }
}

};  // namespace coagulation
