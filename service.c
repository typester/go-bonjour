#import "service.h"
#import "util.h"

@interface BonjourService () {
    void* _ptr;
}
@end

@implementation BonjourService

BonjourService* BonjourServiceNew(void* ptr, char* _domain, char* _type, char* _name, int port) {
    BonjourService* self = [BonjourService new];
    if (self) {
        self->_ptr = ptr;

        NSString* domain = CStringToNSString(_domain);
        NSString* type   =  CStringToNSString(_type);
        NSString* name   = CStringToNSString(_name);

        NSNetService* service = [[NSNetService alloc] initWithDomain:domain
                                                                type:type
                                                                name:name
                                                                port:port];
        self.service = service;
        [service release];

        [name release];
        [type release];
        [domain release];
    }
    return self;
}

BonjourService* BonjourServiceNewFromPtr(void* ptr, void* servicePtr) {
    BonjourService* self = [BonjourService new];
    if (self) {
        self->_ptr = ptr;
        self.service = (NSNetService*)servicePtr;
    }
    return self;
}

void BonjourServiceFree(BonjourService* self) {
    [self release];
}

void BonjourServiceRetain(BonjourService* self) {
    [self retain];
}

char* BonjourServiceGetName(BonjourService* self) {
    return (char*)[self.service.name UTF8String];
}

char* BonjourServiceGetType(BonjourService* self) {
    return (char*)[self.service.type UTF8String];
}

char* BonjourServiceGetDomain(BonjourService* self) {
    return (char*)[self.service.domain UTF8String];
}

char* BonjourServiceGetHostName(BonjourService* self) {
    return (char*)[self.service.hostName UTF8String];
}

int BonjourServiceGetPort(BonjourService* self) {
    return self.service.port;
}

void BonjourServicePublish(BonjourService* self) {
    [self.service publish];
}

void BonjourServiceStop(BonjourService* self) {
    [self.service stop];
}

-(void)dealloc {
    self->_ptr = NULL;
    self.service = nil;
    [super dealloc];
}

@end
