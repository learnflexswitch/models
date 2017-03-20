//
//Copyright [2016] [SnapRoute Inc]
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//	 Unless required by applicable law or agreed to in writing, software
//	 distributed under the License is distributed on an "AS IS" BASIS,
//	 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//	 See the License for the specific language governing permissions and
//	 limitations under the License.
//
// _______  __       __________   ___      _______.____    __    ____  __  .___________.  ______  __    __
// |   ____||  |     |   ____\  \ /  /     /       |\   \  /  \  /   / |  | |           | /      ||  |  |  |
// |  |__   |  |     |  |__   \  V  /     |   (----` \   \/    \/   /  |  | `---|  |----`|  ,----'|  |__|  |
// |   __|  |  |     |   __|   >   <       \   \      \            /   |  |     |  |     |  |     |   __   |
// |  |     |  `----.|  |____ /  .  \  .----)   |      \    /\    /    |  |     |  |     |  `----.|  |  |  |
// |__|     |_______||_______/__/ \__\ |_______/        \__/  \__/     |__|     |__|      \______||__|  |__|
//

package objects

import ()

/*
 * Global Config for ISIS
 */
type IsisGlobal struct {
	baseObj
	Vrf    string `SNAPROUTE: "KEY", ACCESS:"w",  MULTIPLICITY:"1", AUTOCREATE: "true", DEFAULT: "default", DESCRIPTION: "VRF id where ISIS is globally enabled or disabled"`
	Enable bool   `DESCRIPTION: "Global ISIS state in this VRF", DEFAULT: "true"`
}

/*
 * Global State
 */
type IsisGlobalState struct {
	baseObj
	Vrf    string `SNAPROUTE: "KEY", ACCESS:"r",  MULTIPLICITY:"1", DESCRIPTION: "VRF id for which global ISIS state is requested"`
	Enable bool   `DESCRIPTION: "Global ISIS state in this VRF"`
}
