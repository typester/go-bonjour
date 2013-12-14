#import <Foundation/Foundation.h>

@interface BonjourBrowser : NSObject <NSNetServiceBrowserDelegate>

@property (nonatomic, retain) NSNetServiceBrowser* browser;

BonjourBrowser* BonjourBrowserNew(void* ptr);
void BonjourBrowserFree(BonjourBrowser* self);
void BonjourBrowserSearch(BonjourBrowser* self, char* _type, char* _domain);

@end
