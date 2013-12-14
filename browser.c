#import "browser.h"
#import "util.h"

#include <stdint.h>

extern void bonjourDidFoundService(void* ptr, void* servicePtr);
extern void bonjourDidRemoveService(void* ptr, void* servicePtr);
extern void bonjourDidNotSearch(void* ptr, int64_t code);

@interface BonjourBrowser () {
    void* _ptr;
}
@end

@implementation BonjourBrowser

BonjourBrowser* BonjourBrowserNew(void* ptr) {
    BonjourBrowser* self = [BonjourBrowser new];
    if (self) {
        self->_ptr = ptr;
    }
    return self;
}

void BonjourBrowserFree(BonjourBrowser* self) {
    [self release];
}

void BonjourBrowserSearch(BonjourBrowser* self, char* _type, char* _domain) {
    NSString* type = CStringToNSString(_type);
    NSString* domain = CStringToNSString(_domain);

    [self.browser searchForServicesOfType:type inDomain:domain];

    [domain release];
    [type release];
}

-(id)init {
    self = [super init];
    if (self) {
        NSNetServiceBrowser* browser = [NSNetServiceBrowser new];
        self.browser = browser;
        [browser release];
        self.browser.delegate = self;
    }
    return self;
}

-(void)dealloc {
    self->_ptr = NULL;
    self.browser = nil;
    [super dealloc];
}

-(void)netServiceBrowser:(NSNetServiceBrowser *)netServiceBrowser didFindService:(NSNetService *)netService moreComing:(BOOL)moreServicesComing {
    bonjourDidFoundService(self->_ptr, (void*)netService);
}

-(void)netServiceBrowser:(NSNetServiceBrowser *)netServiceBrowser didRemoveService:(NSNetService *)netService moreComing:(BOOL)moreServicesComing {
    bonjourDidRemoveService(self->_ptr, (void*)netService);
}

-(void)netServiceBrowser:(NSNetServiceBrowser *)netServiceBrowser didNotSearch:(NSDictionary *)errorInfo {
    int64_t code = [(NSNumber*)errorInfo[NSNetServicesErrorCode] longLongValue];
    bonjourDidNotSearch(self->_ptr, code);
}

@end
