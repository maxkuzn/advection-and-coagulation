cmake_minimum_required(VERSION 3.8)
project(advection-and-coagulation)


if(CMAKE_CXX_COMPILER_ID MATCHES "Clang")
  set(CLANG 1)
endif()

if (ASAN)
    if(NOT CLANG)
        message(FATAL_ERROR "Cannot enable ASAN unless using Clang")
    endif()
    message("Building with address sanitizer")
    set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -fsanitize=address -fsanitize-address-use-after-scope -fno-omit-frame-pointer")
    set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -fsanitize=address -fsanitize-address-use-after-scope -fno-omit-frame-pointer")
    set(CMAKE_EXE_LINKER_FLAGS "${CMAKE_EXE_LINKER_FLAGS} -fsanitize=address")
    add_definitions(-DNO_TIMEOUT)
endif()

if (TSAN)
    if(NOT CLANG)
        message(FATAL_ERROR "Cannot enable TSAN unless using Clang")
    endif()
    message("Building with thread sanitizer")
    set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -fsanitize=thread")
    set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -fsanitize=thread")
    set(CMAKE_EXE_LINKER_FLAGS "${CMAKE_EXE_LINKER_FLAGS} -fsanitize=thread")
    add_definitions(-DNO_TIMEOUT)
endif()

set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -Wall -Wextra -Werror")
set(CMAKE_EXPORT_COMPILE_COMMANDS  ON)
set(CMAKE_CXX_STANDARD             17)

set(CMAKE_C_ARCHIVE_CREATE   "<CMAKE_AR> Scr <TARGET> <LINK_FLAGS> <OBJECTS>")
set(CMAKE_CXX_ARCHIVE_CREATE "<CMAKE_AR> Scr <TARGET> <LINK_FLAGS> <OBJECTS>")
set(CMAKE_C_ARCHIVE_FINISH   "<CMAKE_RANLIB> -no_warning_for_no_symbols -c <TARGET>")
set(CMAKE_CXX_ARCHIVE_FINISH "<CMAKE_RANLIB> -no_warning_for_no_symbols -c <TARGET>")

enable_testing()
set( CMAKE_RUNTIME_OUTPUT_DIRECTORY ${CMAKE_BINARY_DIR}/bin )

add_subdirectory(library)

add_executable(main main.cpp)
target_link_libraries(main library)
