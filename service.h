#import <Foundation/Foundation.h>

@interface BonjourService : NSObject

@property (nonatomic, retain) NSNetService* service;

BonjourService* BonjourServiceNew(void* ptr, char* _domain, char* _type, char* _name, int port);
BonjourService* BonjourServiceNewFromPtr(void* ptr, void* servicePtr);
void BonjourServiceFree(BonjourService* self);

char* BonjourServiceGetName(BonjourService* self);
char* BonjourServiceGetType(BonjourService* self);
char* BonjourServiceGetDomain(BonjourService* self);
char* BonjourServiceGetHostName(BonjourService* self);
int BonjourServiceGetPort(BonjourService* self);

void BonjourServicePublish(BonjourService* self);
void BonjourServiceStop(BonjourService* self);

@end
