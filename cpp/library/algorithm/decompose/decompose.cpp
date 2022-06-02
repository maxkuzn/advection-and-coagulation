#include "decompose.h"

#include <iostream>

namespace decompose {

std::tuple<Eigen::MatrixXd, Eigen::VectorXd, Eigen::MatrixXd> Decompose(Eigen::MatrixXd k) {
    // Eigen::BDCSVD<Eigen::MatrixXd> svd(k, Eigen::ComputeThinV | Eigen::ComputeThinU);
    Eigen::BDCSVD<Eigen::MatrixXd, Eigen::ComputeThinV | Eigen::ComputeThinU> svd(k);

    auto u = svd.matrixU();
    auto s = svd.singularValues();
    auto v = svd.matrixV();

    size_t rank = 0;
    const double eps = 1e-6;
    while (rank < size_t(s.size()) && s[rank] > eps) {
        rank++;
    }

    u = u(Eigen::indexing::all, Eigen::seqN(0, rank));
    s = s(Eigen::seqN(0, rank));
    v = v(Eigen::indexing::all, Eigen::seqN(0, rank));


    std::cout << s << '\n';

    std::printf("(%d, %d) -> (%d, %d) x (%d, %d) x (%d, %d)\n",
                int(k.rows()), int(k.cols()),
                int(u.rows()), int(u.cols()),
                int(s.size()), int(s.size()),
                int(v.cols()), int(v.rows()));

    return std::make_tuple(u, s, v);
}

}
