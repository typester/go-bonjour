#ifndef _util_h_
#define _util_h_

#import <Foundation/Foundation.h>
#include <string.h>

static inline NSString* CStringToNSString(char* c) {
    return [[NSString alloc] initWithBytesNoCopy:c
                                          length:strlen(c)
                                        encoding:NSUTF8StringEncoding
                                    freeWhenDone:YES];
}

#endif
