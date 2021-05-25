import { Component, OnInit } from "@angular/core";
import { OrbitalServerService } from "./orbital/orbital-server.service";

@Component({
	selector: "app-root",
	templateUrl: "./app.component.html",
	styleUrls: ["./app.component.scss"],
})
export class AppComponent implements OnInit {
	title = "client";

	constructor(private orbital: OrbitalServerService) {}

	ngOnInit() {
		this.orbital.test().subscribe();
	}
}
