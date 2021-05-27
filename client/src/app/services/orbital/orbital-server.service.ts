import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { tap } from "rxjs/operators";

@Injectable({
	providedIn: "root",
})
export class OrbitalServerService {
	private static readonly API_LOCATION = "http://localhost:7558";

	private authenticated: boolean;

	constructor(private http: HttpClient) {
		this.authenticated = false;
	}

	isAuthenticated() {
		return this.http.get<AuthenticationState>(this.path("/authenticated"));
	}

	private path(relativePath: string): string {
		return OrbitalServerService.API_LOCATION + relativePath;
	}
}

interface AuthenticationState {
	is_authenticated: boolean;
}
