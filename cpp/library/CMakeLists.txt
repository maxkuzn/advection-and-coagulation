get_filename_component(LIB_INCLUDE_PATH "." ABSOLUTE)
get_filename_component(LIB_PATH "." ABSOLUTE)

file(GLOB_RECURSE LIB_CXX_SOURCES *.cpp)
file(GLOB_RECURSE LIB_HEADERS *.h)

add_library(library STATIC ${LIB_CXX_SOURCES} ${LIB_HEADERS})
target_include_directories(library PUBLIC ${LIB_INCLUDE_PATH})
target_include_directories(library PUBLIC ${LIB_INCLUDE_PATH}/dj_fft)
target_link_libraries(library PUBLIC pthread)
