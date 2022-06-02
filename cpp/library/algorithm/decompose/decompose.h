#pragma once

#include <tuple>
#include <Eigen/Dense>

namespace decompose {

std::tuple<Eigen::MatrixXd, Eigen::VectorXd, Eigen::MatrixXd> Decompose(Eigen::MatrixXd k);

}