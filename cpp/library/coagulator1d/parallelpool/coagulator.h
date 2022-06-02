#pragma once

#include "algorithm/coagulation/coagulator.h"
#include "base/field1d.h"
#include "coagulator1d/coagulator.h"

#include <memory>
#include <vector>
#include <thread>
#include <queue>
#include <tuple>

namespace coagulation {

const size_t number_of_workers = 15;

class ParallelPoolCoagulator1D : public Coagulator1D {
  public:
    explicit ParallelPoolCoagulator1D(std::shared_ptr<Coagulator> base_coagulator, const std::vector<double>& volumes,
                                      size_t batch_size)
            : base_coagulator_(base_coagulator), volumes_(volumes), batch_size_(batch_size) {
        for (size_t i = 0; i < number_of_workers; i++) {
            workers_.emplace_back(&ParallelPoolCoagulator1D::DoWork, this);
        }
    }

    ~ParallelPoolCoagulator1D() override {
        std::unique_lock lock(mtx_);
        stop_ = true;
        cv_.notify_all();

        lock.unlock();

        for (auto&& w: workers_) {
            w.join();
        }
    }

    std::pair<Field1D*, Field1D*> Process(Field1D* field, Field1D* buff) override;

  private:
    void DoWork();

    std::shared_ptr<Coagulator> base_coagulator_;
    std::vector<double> volumes_;
    const size_t batch_size_;

    std::vector<std::thread> workers_;

    std::mutex mtx_;
    std::condition_variable cv_;
    bool stop_;
    std::queue<std::tuple<Field1D*, Field1D*, size_t, size_t>> queue_;

};

}  // namespace coagulation
