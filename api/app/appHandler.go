package app

// AppInstalled: 앱 설치
// 발행한 앱이 특정 스토어에 설치된 경우, 앱 설치 이벤트(app-installed)가 전달됩니다.
//
//	{
//		"event": "app-installed",       // 이벤트 구분자
//		"store": "<store-id>",          // 대상 스토어 ID
//		"accessKey": "<app-access-key>" // 앱 엑세스 키
//	}
//
//	앱 서버는 해당 이벤트를 활용하여 다음과 같은 작업을 수행할 수 있습니다.
//
// - 스토어 별로 앱이 설치된 경우 해당 스토어를 대행하여 앱으로서 인증할 수 있는 앱 엑세스 키를 전달받고 저장
// - 스토어 별로 앱이 설치된 경우 앱의 동작상 필요한 디폴트 설정들을 미리 준비하는 작업 (앱의 동작을 위한 최소 시스템적 조건을 셋업)
func AppInstalled() {

	// storeId와 accessKey를 DB에 저장/ 일단 저장만

}

// AppDeactived: 앱 비활성화
// 특정 스토어에 앱이 설치된 후 비활성화된 경우, 앱 비활성화 이벤트(app-deactivated)가 전달됩니다.
//
//	{
//		"event": "app-deactivated",     // 이벤트 구분자
//		"store": "<store-id>"           // 대상 스토어 ID
//	}
//
// 앱에서 스토어를 위해 제공하던 기능을 필요시 중단(Pause)시킬 수 있습니다. (CRON, Scheduler 등)
func AppDeactived() {

}

// AppReactive: 앱 재활성화
// 특정 스토어에서 앱이 재활성화된 경우, 앱 재활성화 이벤트(app-reactivated)가 전달됩니다.
//
//	{
//		"event": "app-reactivated",     // 이벤트 구분자
//		"store": "<store-id>"           // 대상 스토어 ID
//	}
//
// 중단된 기능을 다시 재개(Resume)시킬 수 있습니다. (CRON, Scheduler 등)
func AppReactive() {

}

// Uninstalled: 앱 제거
// 특정 스토어에서 앱이 제거된 경우, 앱 제거 이벤트(app-uninstalled)가 전달됩니다.
// {
// 	"event": "app-uninstalled",     // 이벤트 구분자
// 	"store": "<store-id>"           // 대상 스토어 ID
// }
// 앱 서버는 해당 이벤트를 활용하여 다음과 같은 작업을 수행할 수 있습니다.

// - 최초 앱이 설치된 경우 전달받은 앱 엑세스 키를 DB/스토리지에서 제거할 수 있습니다.
// - 앱이 해당 스토어를 위해 저장해놓은 정보를 모두 제거(클린업)할 수 있습니다.
func AppUninstalled() {

}
