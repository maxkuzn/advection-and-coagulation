#include "progress.h"

#include <sstream>

std::string progress(size_t performed, size_t total) {
    constexpr size_t kBarWidth = 60;

    std::stringstream out;
    out << '[';
    double perc = static_cast<double>(performed) / total;
    size_t pos = perc * kBarWidth;
    for (size_t i = 0; i < kBarWidth; i++) {
        if (i < pos) {
            out << '=';
        } else if (i == pos) {
            out << '>';
        } else {
            out << ' ';
        }
    }
    out << ']';

    out << ' ' << static_cast<int>(100 * perc) << '%';
    out << ' ' << performed << '/' << total;

    if (performed == total) {
        out << '\n';
    }

    return out.str();
}
