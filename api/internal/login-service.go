/* The AGPLv3 License (AGPLv3)

Copyright (c) 2022 Zhao Zhenhua <zhao.zhenhua@gmail.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as
published by the Free Software Foundation, either version 3 of the
License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>. */

package internal

type LoginService interface {
	Login(vaultID string, nonce string, signature string) bool
}

type loginService struct {
	VaultID   string
	Nonce     string
	Signature string
}

// if sign by vaultID's controller, return true
func (service *loginService) Login(vaultID string, nonce string, signature string) bool {
	return true
}

func NewLoginService() LoginService {
	return &loginService{}
}
