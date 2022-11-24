#include <iostream>

#include <webp/encode.h>
#include <webp/decode.h>

#ifdef __cplusplus
extern "C"
{
#endif
    uint8_t *allocateBuffer(size_t len)
    {
        uint8_t *buffer = new uint8_t[len];
        return buffer;
    }

    void dump(uint8_t *arr, size_t len)
    {
        size_t i{0};
        std::cout << "You callled dump! size: " << len << "\n";
        for (; i < len; i++)
        {
            std::cout << int(arr[i]) << " ";
        }

        std::cout
            << "\n"
            << "as 4 byte int: " << (size_t *)arr
            << "\n"
            << std::flush;
    }

    // since we're currently not successful allocating memory from javascript to be used in C, we
    // provide allocation inside here
    uint8_t *alloc(size_t size)
    {
        uint8_t *ref = (uint8_t *)malloc(size);
        return ref;
    }

    void webpFree(uint8_t *mem)
    {
        WebPFree(mem);
    }

    void falloc(uint8_t *mem)
    {
        free(mem);
    }

    // https://web.dev/emscripting-a-c-library/
    uint8_t *encodeRGBA(uint8_t *rgba, int width, int height, int stride, float qualityFactor, uint8_t *len)
    {
        uint8_t *img_out;
        size_t size;

        std::cout << "STARTING TO ENCODE\n";
        size = WebPEncodeRGBA(rgba, width, height, stride, qualityFactor, &img_out);

        std::cout << "Size is: "
                  << ((int)size) << ". \n";

        *len = size;

        return img_out;
    }

#ifdef __cplusplus
}
#endif

int main()
{
    uint8_t *gg = allocateBuffer(10);
    gg[0] = 1;
    gg[1] = 2;
    gg[2] = 3;
    std::cout << "Hello webp\n"
              << int(gg[0]) << " "
              << "\n";
}
